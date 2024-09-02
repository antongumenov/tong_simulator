package controller

import (
	"sync"
	"sync/atomic"
	"time"
)

// интерфейсы
type Repository interface {
	GetNext() int64
	Reset()
}

type Producer interface {
	Subscrube(Consumer)
	UnSubscribe(Consumer)
	NotifyListeners()
}

type Consumer interface {
	Update()
	GetName() string
}

// вью модель
type Controller struct {
	repo      Repository
	consumers map[string]Consumer
	sync.RWMutex
	dumpLevel int64
	loadCell  int64
	encoder   int64
	isRun     atomic.Bool // запущено вращение
	isTorque  atomic.Bool // пошел момент
	isDumpOn  atomic.Bool // сработал дамп
	isMax     atomic.Bool // достиг максимума
}

// констурктор
func New(repo Repository) *Controller {
	return &Controller{
		repo:      repo,
		consumers: make(map[string]Consumer, 10),
		dumpLevel: 9000,
		loadCell:  0,
		encoder:   0,
	}
}

// реалтзация продюссера
func (ctrl *Controller) Subscrube(c Consumer) {
	ctrl.Lock()
	ctrl.consumers[c.GetName()] = c
	ctrl.Unlock()
}

func (ctrl *Controller) UnSubscribe(c Consumer) {
	ctrl.Lock()
	delete(ctrl.consumers, c.GetName())
	ctrl.Unlock()
}

func (ctrl *Controller) NotifyListeners() {
	ctrl.RLock()
	for _, c := range ctrl.consumers {
		c.Update()
	}
	ctrl.RUnlock()
}

// Геттеры
func (ctrl *Controller) GetLoadCell() int64 {
	return atomic.LoadInt64(&ctrl.loadCell)
}

func (ctrl *Controller) GetEncoder() int64 {
	return atomic.LoadInt64(&ctrl.encoder)
}

func (ctrl *Controller) GetDumpLevel() int64 {
	return atomic.LoadInt64(&ctrl.dumpLevel)
}

// Сеттеры
func (ctrl *Controller) EncoderAddOne() {
	atomic.StoreInt64(&ctrl.encoder, ctrl.GetEncoder()+1)
}

func (ctrl *Controller) SetDumpLevel(level int64) {
	atomic.StoreInt64(&ctrl.dumpLevel, int64(level))
}

// обработчики
func (ctrl *Controller) Dump() {
	// когда идет вращение останавливает вращение
	if ctrl.isRun.Load() {
		ctrl.isRun.Store(false)
	}
}

func (ctrl *Controller) Reset() {
	// только когда идет вращение, нет момента, сбрасывает обороты
	if ctrl.isRun.Load() && !ctrl.isTorque.Load() {
		atomic.StoreInt64(&ctrl.encoder, 0)
	}
	// только когда вражение остановлено сбрасывает все
	if !ctrl.isRun.Load() {
		atomic.StoreInt64(&ctrl.encoder, 0)
		atomic.StoreInt64(&ctrl.loadCell, 0)
		ctrl.isTorque.Store(false)
		ctrl.isMax.Store(false)
		ctrl.isDumpOn.Store(false)
		ctrl.repo.Reset()
		ctrl.NotifyListeners()
	}
}

func (ctrl *Controller) TorqueOn() {
	// включает вращение момент когда идет вращение
	if ctrl.isRun.Load() {
		ctrl.isTorque.Store(true)
	}
}

func (ctrl *Controller) Rotate() {
	// начинает вращение только когда не запущено, нет максимума
	if !ctrl.isRun.Load() && !ctrl.isMax.Load() {
		go func() {
			// запускаю вращение
			ctrl.isRun.Store(true)
			for {
				// начинаю вращение
				// продолжаю только если запущено вращение
				if !ctrl.isRun.Load() {
					break
				}

				// увеличиваю обороты
				ctrl.EncoderAddOne()

				// если нажали генерацию момента
				if ctrl.isTorque.Load() {
					atomic.StoreInt64(&ctrl.loadCell, ctrl.repo.GetNext())

				}

				// уведомляю подписчиков
				ctrl.NotifyListeners()

				// если момент привысил отсечку, выхожу
				if atomic.LoadInt64(&ctrl.loadCell) >= atomic.LoadInt64(&ctrl.dumpLevel) {
					ctrl.isRun.Store(false)
					ctrl.isMax.Store(true)
					break
				}

				// будем подождать
				time.Sleep(time.Millisecond)
			}
		}()
	}
}
