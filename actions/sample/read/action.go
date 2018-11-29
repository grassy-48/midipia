package read

import (
	"bufio"
	"fmt"
	"os"

	"github.com/algoGuy/EasyMIDI/smfio"
	cli "gopkg.in/urfave/cli.v1"
)

// Read midi file
func Read(c *cli.Context) error {
	// Open test midi file
	file, _ := os.Open("data/outputMidi.mid")
	defer file.Close()

	// Read and save midi to smf.MIDIFile struct
	midi, err := smfio.Read(bufio.NewReader(file))

	if err != nil {
		return err
	}

	// Get zero track
	track := midi.GetTrack(0)

	// Print number of midi tracks
	fmt.Println(midi.GetTracksNum())

	// Get all midi events via iterator
	iter := track.GetIterator()

	for iter.MoveNext() {
		fmt.Println(iter.GetValue())
	}
	return nil
}
