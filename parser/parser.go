package parser

import (
    "archive/zip"
    "errors"
    "io"
    "io/ioutil"
    "os"
)

func readZipFile(reader io.ReaderAt, size int64) ([]byte, error) {
    r, err := zip.NewReader(reader, size)
    if err != nil {
        return nil, err
    }

    for _, f := range r.File {
        if f.Name != "outfrom.xml" {
            continue
        }

        rc, err := f.Open()
        if err != nil {
            return nil, err
        }

        content, err := ioutil.ReadAll(rc)
        if err != nil {
            return nil, err
        }

        rc.Close()
        return content, nil
    }

    return nil, errors.New("outfrom.xml not found")
}

func Parse(reader io.ReaderAt, size int64) (*Fast, error) {
    content, err := readZipFile(reader, size)
    if err != nil {
        return nil, err
    }

    return ParseXML(content)
}

func ParseFile(filename string) (*Fast, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }

    defer file.Close()

    fileinfo, err := file.Stat()
    if err != nil {
        return nil, err
    }

    return Parse(file, fileinfo.Size())
}
