using MainApp.Models;

namespace MainApp.Helpers
{
    public class DateTypeHelper
    {
        public static object degree;
        public static ShelfDegree GetDateType(string dateStr)
        {
            if ((dateStr.Contains("дн")) || (dateStr.Contains("де")))
            {
                return ShelfDegree.Day;
            }
            else if (dateStr.Contains("ме"))
            {
                return ShelfDegree.Month;
            }
            else
            {
                return ShelfDegree.Year;
            }
        }
    }
}
