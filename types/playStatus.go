package types

type Parent struct {
  Content Content `json:"content"`
}

type Content struct {
  IsPlaying bool `json:"is_playing"`
  device map[string]interface{} `json:"device"`
  shuffleState bool `json:"shuffle_state"`
  repeatState string `json:"repeat_state"`
  timestamp int64 `json:"timestamp"`
  progressMs int32 `json:"progress_ms"`
  item map[string]interface{} `json:"item"`
  currentlyPlayingType string `json:"currently_playing_type"`
  action map[string]interface{} `json:"actions"`
}
