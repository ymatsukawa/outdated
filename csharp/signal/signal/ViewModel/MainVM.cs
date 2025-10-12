using System.ComponentModel;
using System.Runtime.CompilerServices;
using System.Windows.Input;
using signal.UI.Picker;
using signal_cli.http;
using signal_cli.io;

namespace signal.ViewModel
{
    public class MainVM : INotifyPropertyChanged
    {
        private string method = string.Empty;
        private string url = string.Empty;
        private string header = string.Empty;
        private string body = string.Empty;

        private string resHeader = string.Empty;
        private string resBody = string.Empty;

        private bool isRunningProcess = false;
        private string savePath = string.Empty;

        public ICommand CommandRequest { get; private set; }
        public ICommand CommandCopyRes { get; private set; }
        public ICommand CommandSave { get; private set; }
        public ICommand CommandLoad { get; private set; }

        public MainVM()
        {
            CommandRequest = new Command(
                execute: async () =>
                {
                    var builder = new RequestBuilder();
                    var request = builder.Method(Method).Url(Url).Header(Header).Body(Body).Build();
                    
                    await request.Run();
                    ResHeader = request.ResponseHeader;
                    ResBody = request.ResponseBody;
                });


            CommandCopyRes = new Command<string>(
                execute: async (string copyType) =>
                {
                    if(copyType.Equals("header") && !string.IsNullOrEmpty(ResHeader))
                    {
                        await Clipboard.Default.SetTextAsync(ResHeader);
                    } else if(copyType.Equals("body") && !string.IsNullOrEmpty(ResBody))
                    {
                        await Clipboard.Default.SetTextAsync(ResBody);
                    }
                });


            CommandSave = new Command(
                execute: async () =>
                {
                    try
                    {
                        Page? page = ViewHelper.MainPage();
                        if (string.IsNullOrEmpty(SavePath))
                        {
                            SavePath = await page.DisplayPromptAsync("Save Folder Path", "Where is 'directory' to save file?");
                        }
                        if (!Path.Exists(SavePath))
                        {
                            await page.DisplayAlert("Alert", string.Format("save path is '{0}' and not found. Set again.", this.SavePath), "OK");
                            return;
                        }
                        RequestSave.Save(SavePath, Method, Url, Header, Body);
                    }
                    catch(Exception ex)
                    {

                    }
                });


            CommandLoad = new Command(
                execute: async () =>
                {
                    string[] formats = [".json"];
                    string? filePath = await FilePick.PickUp(formats);
                    if (filePath == null)
                    {
                        return;
                    }

                    var r = RequestLoad.Load(filePath);
                    if (r == null)
                    {
                        return;
                    }
                    Method = r.Method;
                    Url = r.Url;
                    Header = r.Header;
                    Body = r.Body;
                });

        }

        public string Method
        {
            set { SetProperty(ref method, value); }
            get { return method; }
        }
        public string Url
        {
            set { SetProperty(ref url, value); }
            get { return url; }
        }
        public string Header
        {
            set { SetProperty(ref header, value); }
            get { return header; }
        }
        public string Body
        {
            set { SetProperty(ref body, value); }
            get { return body; }
        }

        public string ResHeader
        {
            set { SetProperty(ref resHeader, value); }
            get { return resHeader; }
        }
        public string ResBody
        {
            set { SetProperty(ref resBody, value); }
            get { return resBody; }
        }

        public bool IsRunningProcess
        {
            set { SetProperty(ref isRunningProcess, value); }
            get { return isRunningProcess; }
        }
        public string SavePath
        {
            set { SetProperty(ref savePath, value); }
            get { return savePath; }
        }


        bool SetProperty<T>(ref T storage, T value, [CallerMemberName] string propertyName = null)
        {
            if (Equals(storage, value))
            {
                return false;
            }

            storage = value;
            OnPropertyChanged(propertyName);
            return true;
        }

        public event PropertyChangedEventHandler PropertyChanged;
        protected void OnPropertyChanged([CallerMemberName] string propertyName = null)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(propertyName));
        }
    }
}
