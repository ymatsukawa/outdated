using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace signal_cli.http
{
    public interface IRequestBuilder
    {
        Request Build();
        RequestBuilder Method(string method);
        RequestBuilder Url(string url);
        RequestBuilder Header(string header);
        RequestBuilder Body(string body);
        RequestBuilder TimeOut(int timeout);
        RequestBuilder ProxyUrl(string url);
        RequestBuilder ProxyId(string id);
        RequestBuilder ProxyPass(string pass);
    }

    public class RequestBuilder : IRequestBuilder
    {
        private Request Request;

        public RequestBuilder()
        {
            this.Request = new Request();
        }

        public Request Build()
        {
            return this.Request;
        }

        public RequestBuilder Method(string method)
        {
            this.Request.Method = method;
            return this;
        }

        public RequestBuilder Url(string url)
        {
            this.Request.Url = url;
            return this;
        }

        public RequestBuilder Header(string header)
        {
            this.Request.Header = header;
            return this;
        }

        public RequestBuilder Body(string body)
        {
            this.Request.Body = body;
            return this;
        }

        public RequestBuilder TimeOut(int timeout)
        {
            this.Request.TimeOut = TimeSpan.FromSeconds(timeout);
            return this;
        }
        public RequestBuilder ProxyUrl(string url)
        {
            this.Request.ProxyUrl = url;
            return this;
        }
        public RequestBuilder ProxyId(string proxyId)
        {
            this.Request.ProxyId = proxyId;
            return this;
        }
        public RequestBuilder ProxyPass(string proxyPass)
        {
            this.Request.ProxyPass = proxyPass;
            return this;
        }
    }
}
