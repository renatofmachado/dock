package services

func redis() *Service {
	service := Service{
		Name:         "Redis",
		Organization: "library",
		Image:        "redis",
		Port:         "6379",
		Template:     "-p ${:port}:6379 -v ${:volume}:/data ${:organization}/${:image_name}:${:tag}",
		Parameters: []Parameter{
			{Key: "volume", Question: "What is the Docker volume name?", DefaultValue: "redis_volume"},
		},
	}

	return &service
}
