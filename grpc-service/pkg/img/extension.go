package img

func MimeToExtension(mimeType string) string {
	mimes := map[string]string{"image/jpeg": "jpeg", "image/jpg": "jpg", "image/png": "png"}
	return mimes[mimeType]
}
