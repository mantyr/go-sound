package test

import (
	s "github.com/padster/go-sound/sounds"
	u "github.com/padster/go-sound/util"
)

// samples.go includes all the example wavs used by the testing package,
// and written .wav golden files to this folder.

func SampleTimedSineSound() s.Sound {
	// Includes: SineWave
	return s.NewTimedSound(s.NewSineWave(261.63), 1000)
}

func SampleTimedSquareSound() s.Sound {
	// Includes: SquareWave
	return s.NewTimedSound(s.NewSquareWave(261.63), 1000)
}

func SampleTimedSawtoothSound() s.Sound {
	// Includes: SawtoothWave
	return s.NewTimedSound(s.NewSawtoothWave(261.63), 1000)
}

func SampleTimedTriangleSound() s.Sound {
	// Includes: SquareWave
	return s.NewTimedSound(s.NewSquareWave(261.63), 1000)
}

func SampleSilence() s.Sound {
	// Includes: TimedSound
	return s.NewTimedSilence(2000.0)
}

func SampleConcat() s.Sound {
	// Includes: TimedSound and MidiToSound
	return s.ConcatSounds(
		s.NewTimedSound(u.MidiToSound(72), 400),
		s.NewTimedSound(u.MidiToSound(74), 400),
		s.NewTimedSound(u.MidiToSound(76), 400),
		s.NewTimedSound(u.MidiToSound(60), 400),
		s.NewTimedSound(u.MidiToSound(67), 1200),
	)
}

func SampleNormalSum() s.Sound {
	// Includes: TimedSound and MidiToSound
	return s.SumSounds(
		s.NewTimedSound(u.MidiToSound(55), 333),
		s.NewTimedSound(u.MidiToSound(59), 333),
		s.NewTimedSound(u.MidiToSound(62), 333),
		s.NewTimedSound(u.MidiToSound(65), 333),
		s.NewTimedSound(u.MidiToSound(67), 333),
	)
}

func SampleMultiply() s.Sound {
	// Includes: TimedSound and SineWave
	all := make([]s.Sound, 20)
	for i := 0; i < len(all); i++ {
		all[i] = s.MultiplyWithClip(s.NewTimedSound(s.NewSineWave(659.25), 200), 0.2+float64(i)/10.0)
	}
	return s.ConcatSounds(all...)
}

func SampleRepeater() s.Sound {
	// Includes: Concat, TimedSound and MidiToSound
	return s.RepeatSound(s.ConcatSounds(
		s.NewTimedSound(u.MidiToSound(50), 400),
		s.NewTimedSound(u.MidiToSound(45), 400),
		s.NewTimedSound(u.MidiToSound(47), 400),
		s.NewTimedSound(u.MidiToSound(42), 400),
		s.NewTimedSound(u.MidiToSound(43), 400),
		s.NewTimedSound(u.MidiToSound(38), 400),
		s.NewTimedSound(u.MidiToSound(43), 400),
		s.NewTimedSound(u.MidiToSound(45), 400),
	), 3)
}

func SampleAdsrEnvelope() s.Sound {
	// Includes: TimedSound and SineWave
	return s.NewADSREnvelope(
		s.NewTimedSound(s.NewSineWave(880.0), 875),
		50, 200, 0.5, 100)
}

func SampleSampler() s.Sound {
	// Includes: TimedSound and SineWave
	return s.LinearSample(s.NewTimedSound(s.NewSineWave(392.00), 500), 2.0)
}

func SampleAddDelay() s.Sound {
	// Includes: Concat, TimedSound and MidiToSound
	return s.AddDelay(s.ConcatSounds(
		s.NewTimedSound(u.MidiToSound(55), 678),
		s.NewTimedSound(u.MidiToSound(59), 678),
		s.NewTimedSound(u.MidiToSound(62), 678),
	), 123)
}

func SampleDenseIIR() s.Sound {
	// Includes: TimedSound and SineWave
	all := make([]s.Sound, 10)
	for i := 0; i < len(all); i++ {
		all[i] = s.NewTimedSound(s.NewSineWave(600*float64(i)/4), 200)
	}
	return s.NewDenseIIR(s.ConcatSounds(all...),
		[]float64{0.8922, -2.677, 2.677, -0.8922},
		[]float64{2.772, -2.57, 0.7961},
	)
}
