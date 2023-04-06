export interface FormField {
  id: string;
  type: 'field';
  columnName: string;
  operator: string;
  value: string | string[] | number | number[];
}

export interface FormGroup {
  id: string;
  type: 'group';
  conjunction: 'and' | 'or';
  children: (FormGroup | FormField)[];
}

export interface FilterForm {
  filterSet: FormGroup;
}
