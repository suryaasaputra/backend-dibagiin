package helpers

import (
	"dibagi/config"
	"fmt"
	"io"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

func UploadToBucket(c *http.Request, key, id string) (string, error) {
	var (
		storageClient *storage.Client
		err           error
	)
	ctx := appengine.NewContext(c)
	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("../secret/dibagi-in-key.json"))
	if err != nil {
		return "", err
	}

	f, _, err := c.FormFile(key)
	if err != nil {
		return "", err
	}

	defer f.Close()
	currentTime := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", currentTime, id)

	sw := storageClient.Bucket(config.BUCKET_NAME).Object(key + "/" + fileName).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		return "", err
	}

	if err := sw.Close(); err != nil {
		return "", err
	}
	photoURL := fmt.Sprintf("%s/%s", config.STORAGE_PATH, sw.Attrs().Name)
	return photoURL, nil
}
