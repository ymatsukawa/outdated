namespace signal_cli.http
{
    internal interface IRequest
    {
        Task Run();
    }

    public class Request : IRequest
    {
        public string ResponseHeader { get; protected set; }
        public string ResponseBody { get; protected set; }

        public string Method { get; set; }
        public string Url { get; set; }
        public string Header { get; set; }
        public string Body { get; set; }

        public TimeSpan TimeOut { get; set; }
        public string ProxyUrl { get; set; }
        public string ProxyId { get; set; }
        public string ProxyPass { get; set; }

        public Request()
        {
            this.ResponseHeader = string.Empty;
            this.ResponseBody = string.Empty;

            this.Method = "GET";
            this.Url = string.Empty;
            this.Header = string.Empty;
            this.Body = string.Empty;

            this.TimeOut = TimeSpan.FromSeconds(10.0d);
            this.ProxyUrl = string.Empty;
            this.ProxyId = string.Empty;
            this.ProxyPass = string.Empty;
        }

        public async Task Run()
        {
            try
            {
                using HttpClient client = new();
                this.SetHttpOption(client);
                IRequestHandler handler = this.Method switch
                {
                    "GET" => new Get(client),
                    "POST" => new Post(client), // ATTENTION: now, application/json only
                    _ => throw new Exception("Unknown method was set"),
                };
                this.ResponseBody = await handler.ExecAsync(this.Url, this.Header, this.Body);
                this.ResponseHeader = handler.ResponseHeader;
            }
            catch(Exception ex)
            {
                this.ResponseBody = $"Error: {ex.Message}";
            }
        }

        private void SetHttpOption(HttpClient client)
        {
            client.Timeout = this.TimeOut;

            if(new string[] { this.ProxyUrl, this.ProxyId, this.ProxyPass }.Any(string.IsNullOrEmpty))
            {
                return;
            }
        }
    }
}
