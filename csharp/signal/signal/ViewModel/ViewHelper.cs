using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace signal.ViewModel
{
    internal class ViewHelper
    {
        public static Page MainPage()
        {
            Page? page = (Application.Current?.MainPage) ?? throw new Exception("page does not exist");
            return page;
        }
    }
}
