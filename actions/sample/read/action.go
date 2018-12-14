package read

import (
	"bufio"
	"fmt"
	"os"

	"github.com/algoGuy/EasyMIDI/smfio"
	cli "gopkg.in/urfave/cli.v1"
)

const defaultFileName = "onkai"

// algoGuy/EasyMID: Example 1
// https://github.com/algoGuy/EasyMIDI#example-1-read-and-get-data-from-midi-file
// Read midi file
func Read(c *cli.Context) error {
	fname := c.String("file")
	if fname == "" {
		fname = defaultFileName
	}
	// Open test midi file
	file, _ := os.Open(fmt.Sprintf("data/%s.mid", fname))
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
