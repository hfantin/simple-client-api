package com.github.hfantin.licensingservice.config

import org.springframework.cloud.client.loadbalancer.LoadBalanced
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.web.client.RestTemplate


@Configuration
class RestTemplateConfig {

    /**
     * this is the restTemplate used by OrganizationRestTemplateClient and must use @LoadBalance annotation
     * to tells springcloud to create a ribbon backed RestTemplate class
     */
    @Bean
    @LoadBalanced
    fun restTemplate() = RestTemplate()


/*
    Its not necessary to use this custom rest template like the book sugests in section 7.10(page 219), 
    maybe its already implemented in latest version of spring security...
    @Primary
    @Bean
    fun getCustomRestTemplate(): RestTemplate {
        val template = RestTemplate()
        val interceptors = template.interceptors
        if (interceptors == null) {
            template.interceptors = listOf<ClientHttpRequestInterceptor>(UserContextInterceptor())
        } else {
            interceptors.add(UserContextInterceptor())
            template.interceptors = interceptors
        }

        return template
    }*/

    /**
     * I dont know if is necessary
     * From page 212 :
     *
     * Finally, all you need to do is modify how the code in the licensing service calls the
     * organization service. You need to ensure that the “Authorization” HTTP header is
     * injected into the application call out to the Organization service. Without Spring Secu-
     * rity, you’d have to write a servlet filter to grab the HTTP header off the incoming licens-
     * ing service call and then manually add it to every outbound service call in the licensing
     * service. Spring OA uth 2 provides a new Rest Template class that supports OA uth 2 calls.
     * The class is called OAuth2RestTemplate . To use the OAuth2RestTemplate class you
     * first need to expose it as a bean that can be auto-wired into a service calling another
     * OA uth 2 protected services. You do this in the licensing-service/src/main/
     * java/com/thoughtmechanix/licenses/Application.java class:
     *
     */
//    @Bean
//    @LoadBalanced
//    fun oauth2RestTemplate(oauth2ClientContext: OAuth2ClientContext, details: OAuth2ProtectedResourceDetails) =
//            OAuth2RestTemplate(details, oauth2ClientContext)

}
