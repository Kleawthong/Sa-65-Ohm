import { ScholarAdminsInterface } from "./IScholarAdmin";
import { ScholarStatusesInterface } from "./IScholarStatus";
import { ScholarTypesInterface } from "./IScholarType";

export interface ScholarshipInterface {
    ID?: number;
    ScholarName: string;
    ScholarAdminID?: number;
    ScholarAdmin?: ScholarAdminsInterface;
    ScholarStatusID?: number;
    ScholarStatus?: ScholarStatusesInterface;
    ScholarTypeID?: number;
    ScholarType?: ScholarTypesInterface;
    ScholarDetail: string;
  }