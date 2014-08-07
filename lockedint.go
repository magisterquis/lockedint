/* lockedint provides a threadsafe int. */
package lockedint

/* TInt represents a threadsafe int. */
type TInt struct {
	m sync.RWMutex
	n int
}

/* New makes a new threadsafe TInt. */
func New() *TInt {
	i := &Int{}
	i.n = 0
	return i
}

/* Add a (possibly negative) value to the TInt. */
func (t *TInt) Add(v int) {
        t.Lock()
        defer t.Unlock()
        t.n += v
}

/* Increment the TInt. */
func (t *TInt) Inc() {
        t.Lock()
        defer t.Unlock()
        t.n++
}

/* Decrement the TInt. */
func (t *Tint) Dec() {
        t.Lock()
        defer t.Unlock()
        t.n--
}

/* Val returns the value of the TInt */
func (t Tint) Val() int {
        t.RLock()
        defer t.RUnlock()
        return t.n
}
