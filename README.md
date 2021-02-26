# LicenseInjector
Inject the license into the binary

``` golang
import "github.com/micropkg/LicenseInjector/injector"

var _ = func() int {
    injector.Inject(`
    <Package Name>
    Some License...`)
    return 0
}()
```
