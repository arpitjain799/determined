export const ItemTypes = {
  FIELD: 'field',
  GROUP: 'group',
} as const;

export interface FormField {
  id: string;
  type: typeof ItemTypes.FIELD;
  columnName: string;
  operator: string;
  value: string | string[] | number | number[];
}

export type Conjunction = 'and' | 'or';

export interface FormGroup {
  id: string;
  type: typeof ItemTypes.GROUP;
  conjunction: Conjunction;
  children: (FormGroup | FormField)[];
}

export interface FilterFormSet {
  filterSet: FormGroup;
}
