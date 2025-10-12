using signal.ViewModel;
using signal.UI.Picker;


namespace signal
{
    public partial class MainPage : ContentPage
    {
        public MainVM MainVM { get; set; }
        private string SavePath = string.Empty;

        public MainPage()
        {
            InitializeComponent();

            this.MainVM = new MainVM();
            this.BindingContext = this.MainVM;
            this.pickerHttpMethod.SelectedIndex = 0;
        }
    }
}
