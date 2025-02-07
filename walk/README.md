# walk
--
    import "vimagination.zapto.org/python/walk"

Package walk provides a python type walker.

## Usage

#### func  Walk

```go
func Walk(t python.Type, fn Handler) error
```
Walk calls the Handle function on the given interface for each non-nil,
non-Token field of the given R type.

#### type Handler

```go
type Handler interface {
	Handle(python.Type) error
}
```

Handler is used to process python types.

#### type HandlerFunc

```go
type HandlerFunc func(python.Type) error
```

HandlerFunc wraps a func to implement Handler interface.

#### func (HandlerFunc) Handle

```go
func (h HandlerFunc) Handle(t python.Type) error
```
Handle implements the Handler interface.
