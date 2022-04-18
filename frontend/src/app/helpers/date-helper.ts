import {DatePipe} from "@angular/common";

var currentDate = new Date();
export default class DateUtils {


  static getPostDate(datePipe: DatePipe, creationDate: string) {
    const date = datePipe.transform(creationDate, 'MM-dd-yyyy');

    const difference = this.getDifference(new Date(date!), currentDate)
    console.log(difference)
    if (difference == 0) {
      return "Posted today"
    } else if (difference == 1) {
      return "Posted yesterday"
    }
    return "Posted:" + " " + difference + " days ago"
  }

  static getDifference(date1: Date, date2: Date) {
    const Difference_In_Time = date2.getTime() - date1.getTime();
    const ONE_DAY = 1000 * 60 * 60 * 24;
    return Math.floor(Difference_In_Time / ONE_DAY)
  }
}
