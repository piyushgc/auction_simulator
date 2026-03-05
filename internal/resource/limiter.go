package resource

type CPULimiter struct {
	sem chan struct{}
}

func NewCPULimiter(maxCPU int) *CPULimiter {
	return &CPULimiter{
		sem: make(chan struct{}, maxCPU),
	}
}

func (c *CPULimiter) Acquire() {
	c.sem <- struct{}{}
}

func (c *CPULimiter) Release() {
	<-c.sem
}
