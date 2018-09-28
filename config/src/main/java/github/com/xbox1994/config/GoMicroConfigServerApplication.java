package github.com.xbox1994.config;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.SpringCloudApplication;
import org.springframework.cloud.config.server.EnableConfigServer;

@EnableConfigServer
@SpringCloudApplication
public class GoMicroConfigServerApplication {

	public static void main(String[] args) {
		SpringApplication.run(GoMicroConfigServerApplication.class, args);
	}
}
