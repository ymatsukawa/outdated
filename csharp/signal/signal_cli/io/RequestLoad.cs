using System.Text.Json;
using signal_cli.io.serialize.json;

namespace signal_cli.io
{
    public class RequestLoad
    {
        public static Request? Load(string filePath)
        {

            try
            {
                string json = File.ReadAllText(filePath);
                return JsonSerializer.Deserialize<Request>(json);
            }
            catch(Exception ex)
            {
                return null;
            }
        }
    }
}
