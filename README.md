Implements a very simple server that redirects traffic to https. One
application of this (what it was created for) is an ECS service that works
around the following ALB limitation (from 
http://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-load-balancing.html
):

> Note
> 
> Currently, Amazon ECS services can only specify a single load balancer or
> target group. If your service requires access to multiple load balanced ports
> (for example, port 80 and port 443 for an HTTP/HTTPS service), you must use a
> Classic Load Balancer with multiple listeners. To use an Application Load
> Balancer, separate the single HTTP/HTTPS service into two services, where
> each handles requests for different ports. Then, each service could use a
> different target group behind a single Application Load Balancer.

