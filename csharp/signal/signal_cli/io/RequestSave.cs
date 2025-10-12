using System.Text.Json;
using signal_cli.io.serialize.json;

namespace signal_cli.io
{
    public class RequestSave()
    {
        public static void Save(string filePath, string method, string url, string header, string body)
        {
            string savePath = Path.Combine(filePath, string.Format("{0}{1}.json", "request_", DateTime.Now.ToString("yyyyMMdd_hmmss")));
            string json = CreateJson(method, url, header, body);
            try
            {
                using var stream = new StreamWriter(savePath);
                stream.Write(json);
            }
            catch (Exception ex)
            {
                // TODO
            }
        }

        private static string CreateJson(string method, string url, string header, string body)
        {
            // TODO: AssemblyInfo
            var metaInfo = new MetaInfo("0.1");
            var request = new serialize.json.Request
            {
                MetaInfo = metaInfo,
                Method = method,
                Url = url,
                Header = header,
                Body = body
            };
            return JsonSerializer.Serialize(request);
        }
    }
}
