package jp.ymatsukawa;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class ResourceGetterApplication {

	/**
	 * java -jar resource_getter.war -p(arg0) 8080(arg1)
	 * @param args
	 */
	public static void main(String[] args) {
		if(!args[0].equals("-p") || args[1].equals("")) {
			System.getProperties().put("server.port", 8080);			
		} else {
			System.getProperties().put("server.port", args[1]);
		}
		SpringApplication.run(ResourceGetterApplication.class, args);
	}
}
