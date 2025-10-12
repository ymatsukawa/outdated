using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

namespace signal_cli.http
{
    internal interface IRequestHandler
    {
        public string ResponseHeader { get; }
        Task<string> ExecAsync(string url, string header, string body);
    }

    internal partial class HandlerBase
    {
        public string ResponseHeader { get; protected set; } = string.Empty;
        protected HttpClient client;

        public HandlerBase(HttpClient client)
        {
            this.client = client;
        }

        protected void SetHeader(string headers)
        {
            string[] headerLines = headers.Split("\n");
            foreach (string headerLine in headerLines)
            {
                string[] h = headerDivide().Split(headerLine);
                if (h.Length == 2)
                {
                    this.client.DefaultRequestHeaders.Add(h[0], h[1]);
                }
            }
        }

        [GeneratedRegex(@": ")]
        private static partial Regex headerDivide();
    }

    internal partial class Get(HttpClient client) : HandlerBase(client), IRequestHandler
    {
        public async Task<string> ExecAsync(string url, string headers, string body)
        {
            if(!string.IsNullOrEmpty(headers))
            {
                SetHeader(headers);
            }

            var res = await this.client.GetAsync(url);

            this.ResponseHeader = res.Content.Headers.ToString();
            return res.Content.ReadAsStringAsync().Result;
        }
    }

    internal class Post(HttpClient client) : HandlerBase(client), IRequestHandler
    {
        public async Task<string> ExecAsync(string url, string headers, string body)
        {
            if (!string.IsNullOrEmpty(headers))
            {
                SetHeader(headers);
            }

            // WIP: content builder
            var content = new StringContent(body, Encoding.UTF8, "application/json");

            var res = await client.PostAsync(url, content);
            return res.Content.ReadAsStringAsync().Result;
        }
    }
}
