package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"io"
	"os"
)

// HashStringMd5 hashes an MD5 string and return the checksum
func HashStringMd5(hs string) []byte {
	data := []byte(hs)
	hash := md5.Sum(data)
	return hash[:] // Looks like [:] is required to convert array to slice
}

// HashFileMd5 hashes the contents of a file and returns checksum
func HashFileMd5(hs string) []byte {
	file, err := os.Open(hs)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	data := md5.New()
	if _, err := io.Copy(data, file); err != nil {
		panic(err)
	}
	hash := data.Sum(nil)
	return hash[:]
}

// HashStringSha256 hashes a string with SHA256 and returns checksum
func HashStringSha256(hs string) []byte {
	data := []byte(hs)
	hash := sha256.Sum256(data)
	return hash[:]
}

// HashFileSha256 hashes the contents of a file and returns the checksum
func HashFileSha256(hs string) []byte {
	file, err := os.Open(hs)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	data := sha256.New()
	if _, err := io.Copy(data, file); err != nil {
		panic(err)
	}
	hash := data.Sum(nil)
	return hash[:]
}

// HashStringSha1 hashes a string with SHA1 and returns checksum
func HashStringSha1(hs string) []byte {
	data := []byte(hs)
	hash := sha1.Sum(data)
	return hash[:]
}

// HashFileSha1 hashes a string with SHA1 and returns checksum
func HashFileSha1(hs string) []byte {
	file, err := os.Open(hs)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	data := sha1.New()
	if _, err := io.Copy(data, file); err != nil {
		panic(err)
	}

	hash := data.Sum(nil)
	return hash[:]
}
