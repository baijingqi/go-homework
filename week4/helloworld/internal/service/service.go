package service

import (
    "github.com/google/wire"
    service "helloworld/internal/service/praise"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService,service.NewPraiseService())
