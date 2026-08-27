package channelstream

import "errors"

var ErrSource = errors.New("source interrupted")

func Stream(items []string, failAt int) (<-chan string, <-chan error) {
    data := make(chan string)
    errs := make(chan error)
    go func() {
        for index, item := range items {
            if index == failAt { errs <- ErrSource; return }
            data <- item
        }
    }()
    return data, errs
}
