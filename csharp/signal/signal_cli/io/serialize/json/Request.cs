namespace signal_cli.io.serialize.json
{
    public class Request
    {
        public MetaInfo? MetaInfo { get; set; } = null; 
        public string Method { get; set; } = string.Empty;
        public string Url { get; set; } = string.Empty;
        public string Header { get; set; } = string.Empty;
        public string Body { get; set; } = string.Empty;
    }
}
