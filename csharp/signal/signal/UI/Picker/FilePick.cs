namespace signal.UI.Picker
{
    public class FilePick
    { 
        public static async Task<string?> PickUp(string[] requiredFormts, string title = "select file")
        {
            var opts = new PickOptions
            {
                PickerTitle = title,
                FileTypes = new FilePickerFileType(new Dictionary<DevicePlatform, IEnumerable<string>>
                    {
                        { DevicePlatform.WinUI, requiredFormts },
                        { DevicePlatform.macOS, requiredFormts },
                    })
            };
            var file = await FilePicker.Default.PickAsync(opts);

            return file?.FullPath;
        }
    }
}
