package samjung

import (
	"encoding/gob"
	"os"
	"fmt"
)

func (r *Samjung) makeIndex(pk, offset uint64) error {
	// add index in-memory
	r.indexMap[pk] = offset
	
	// write map to file
	err := r.writeIndexToFile()
	if err != nil {
		return nil
	}
	
	return nil
}


func (r *Samjung) writeIndexToFile() error {
	f, err := os.OpenFile(r.baseDir+"/"+indexFile, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer f.Close()
		
	// Create an encoder and send a value.
	enc := gob.NewEncoder(f)
	err = enc.Encode(&r.indexMap)
	if err != nil {
		return err
	}

	return nil
}

func (r *Samjung) readIndexFromFile() error {
	f, err := os.Open(r.baseDir+"/"+indexFile)
	if err != nil {
		return err
	}
	defer f.Close()
	
	// Create a decoder and receive a value.
	dec := gob.NewDecoder(f)
	err = dec.Decode(&r.indexMap)
	if err != nil {
		return err
	}
	
	fmt.Printf("%v", r.indexMap)
	return nil
}