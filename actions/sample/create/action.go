package create

import (
	"bufio"
	"os"

	"github.com/algoGuy/EasyMIDI/smf"
	"github.com/algoGuy/EasyMIDI/smfio"
	cli "gopkg.in/urfave/cli.v1"
)

// Create midi file
func Create(c *cli.Context) error {
	// Create division
	division, err := smf.NewDivision(960, smf.NOSMTPE)
	if err != nil {
		return err
	}

	// Create new midi struct
	midi, err := smf.NewSMF(smf.Format0, *division)
	if err != nil {
		return err
	}

	// Create track struct
	track := &smf.Track{}

	// Add track to new midi struct
	err = midi.AddTrack(track)
	if err != nil {
		return err
	}

	// Create some midi and meta events
	midiEventOne, err := smf.NewMIDIEvent(0, smf.NoteOnStatus, 0x00, 0x30, 0x50)
	if err != nil {
		return err
	}

	midiEventTwo, err := smf.NewMIDIEvent(10000, smf.NoteOnStatus, 0x00, 0x30, 0x00)
	if err != nil {
		return err
	}

	metaEventOne, err := smf.NewMetaEvent(0, smf.MetaEndOfTrack, []byte{})
	if err != nil {
		return err
	}

	// Add created events to track
	if err := track.AddEvent(midiEventOne); err != nil {
		return err
	}
	if err := track.AddEvent(midiEventTwo); err != nil {
		return err
	}
	if err := track.AddEvent(metaEventOne); err != nil {
		return err
	}

	// Save to new midi source file
	outputMidi, err := os.Create("data/outputMidi.mid")
	if err != nil {
		return err
	}
	defer outputMidi.Close()

	// Create buffering stream
	writer := bufio.NewWriter(outputMidi)
	smfio.Write(writer, midi)
	return writer.Flush()
}
