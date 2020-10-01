package services

func mongo() *Service {
	service := Service{
		Name:         "mongo",
		Organization: "library",
		Image:        "redis",
		Port:         "27017",
		Template:     "-p ${:port}:27017 -e MONGO_INITDB_ROOT_USERNAME=${:root_user} -e MONGO_INITDB_ROOT_PASSWORD=${:root_password} -v ${:volume}:/data/db ${:organization}/${:image_name}:${:tag}",
		Parameters: []Parameter{
			{Key: "volume", Question: "What is the Docker volume name?", DefaultValue: "mongo_data"},
			{Key: "root_user", Question: "What is the root username?", DefaultValue: "admin"},
			{Key: "root_password", Question: "What is the root password?", DefaultValue: "password"},
		},
	}

	return &service
}
