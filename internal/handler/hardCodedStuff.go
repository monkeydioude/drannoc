package handler

import "time"

// this file should be gradually removed

const tokenDuration = 10 * time.Second
const tokenRemakeThreshold = (tokenDuration * 8) / 10
