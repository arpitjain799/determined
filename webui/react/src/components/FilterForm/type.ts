import { ValueOf } from 'shared/types';

export const FormType = {
  Field: 'field',
  Group: 'group',
} as const;

export type FormType = ValueOf<typeof FormType>;

export type FormField = {
  readonly id: string;
  readonly type: typeof FormType.Field;
  columnName: string;
  operator: Operator;
  value: string | string[] | number | number[] | undefined;
};

export const Conjunction = {
  and: 'and',
  or: 'or',
} as const;

export type Conjunction = ValueOf<typeof Conjunction>;

export type FormGroup = {
  readonly id: string;
  readonly type: typeof FormType.Group;
  conjunction: Conjunction;
  children: (FormGroup | FormField)[];
};

export type FilterFormSet = {
  filterGroup: FormGroup;
};

export type Operator =
  | 'contains'
  | 'in'
  | 'is'
  | 'eq'
  | 'greater'
  | 'greaterEq'
  | 'isEmpty'
  | 'isNot'
  | 'less'
  | 'lessEq'
  | 'notContain'
  | 'notEmpty'
  | 'notEq'
  | 'notIn';

export const OperatorMap: Record<Operator, string> = {
  contains: 'contains',
  eq: '=',
  greater: '>',
  greaterEq: '>=',
  in: 'in',
  is: 'is',
  isEmpty: 'is empty',
  isNot: 'is not',
  less: '<',
  lessEq: '<=',
  notContain: 'not contains',
  notEmpty: 'not empty',
  notEq: '!=',
  notIn: 'not in',
} as const;
