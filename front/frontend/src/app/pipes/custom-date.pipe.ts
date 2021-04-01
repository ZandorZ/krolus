import { DatePipe } from "@angular/common";
import { Pipe, PipeTransform } from "@angular/core";

@Pipe({
    name: "customDate"
})
export class CustomDatePipe extends DatePipe implements PipeTransform {
    transform(value: any, args?: any): any {
        const d2 = new Date(value);
        const d1 = new Date();

        //same day
        if (
            d1.getFullYear() === d2.getFullYear() &&
            d1.getMonth() === d2.getMonth() &&
            d1.getDate() === d2.getDate()
        ) {
            return super.transform(d2, "HH:mm");
        }

        // //same year
        if (d1.getFullYear() === d2.getFullYear()) {
            return super.transform(d2, "MMM d");
        }

        return super.transform(d2, "MMM d, YY");
    }
}