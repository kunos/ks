package ks

type GameTime struct {
	DeltaT       float32
	SmoothDeltaT float32
	Now          float64
}

func (gt *GameTime) Update() {

	now := float64(GetTimeMillis())

	elapsed := now - gt.Now
	gt.DeltaT = float32(elapsed * 0.001)
	gt.Now = now

	gt.SmoothDeltaT = gt.SmoothDeltaT*0.9 + gt.DeltaT*0.1

}
