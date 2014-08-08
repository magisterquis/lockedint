/* lockedint provides a threadsafe int. */
package lockedint

import (
        "sync"
)

/* TInt represents a threadsafe int. */
type TInt struct {
	m sync.RWMutex
	n int
}

/* New makes a new threadsafe TInt. */
func New() *TInt {
	i := &TInt{}
	i.n = 0
	return i
}

/* Add a (possibly negative) value to the TInt. */
func (t *TInt) Add(v int) int {
        t.m.Lock()
        defer t.m.Unlock()
        t.n += v
        return t.n
}

/* Increment the TInt. */
func (t *TInt) Inc() int {
        t.m.Lock()
        defer t.m.Unlock()
        t.n++
        return t.n
}

/* Decrement the TInt. */
func (t *TInt) Dec() int {
        t.m.Lock()
        defer t.m.Unlock()
        t.n--
        return t.n
}

/* Val returns the value of the TInt */
func (t TInt) Val() int {
        t.m.RLock()
        defer t.m.RUnlock()
        return t.n
}
