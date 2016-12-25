package gosshtool

import (
	"github.com/golang/crypto/ssh"
)

type Tunnel struct {
	Client *ssh.Client
}
