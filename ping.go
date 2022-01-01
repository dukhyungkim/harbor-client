package harbor

func (c *client) Ping() (string, error) {
	data, err := c.getText(urlPing, false)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
