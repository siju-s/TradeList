import {Injectable} from "@angular/core";
import {Subject} from "rxjs";

@Injectable()
export class DataService {
  public categories: Subject<any> = new Subject();
  public subcategories: Subject<any> = new Subject();

  constructor() {}
}
