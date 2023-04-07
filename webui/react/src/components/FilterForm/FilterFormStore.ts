import { Observable, observable } from 'micro-observables';

import { Conjunction, FilterFormSet, FormField, FormGroup, FormType, Operator } from './type';

const INIT_FORMSET: FilterFormSet = {
  filterGroup: { children: [], conjunction: Conjunction.and, id: 'ROOT', type: FormType.Group }, // default
};

const getInitGroup = (): FormGroup => {
  return {
    children: [],
    conjunction: Conjunction.and,
    id: crypto.randomUUID(),
    type: FormType.Group,
  };
};

const getInitField = (): FormField => {
  return {
    columnName: 'id',
    id: crypto.randomUUID(),
    operator: 'contains',
    type: FormType.Field,
    value: undefined,
  };
};

export class FormClassStore {
  #formset = observable<FilterFormSet>(INIT_FORMSET);

  constructor(data?: FilterFormSet) {
    if (data) {
      this.#formset = observable<FilterFormSet>(data);
    }
  }

  public get formset(): Observable<Readonly<FilterFormSet>> {
    return this.#formset.readOnly();
  }

  public setFieldValue(
    id: string,
    keyType:
      | keyof Pick<FormField, 'columnName' | 'operator' | 'value'>
      | keyof Pick<FormGroup, 'conjunction'>,
    value: string | string[] | number | number[],
  ): void {
    const set = this.#formset.get().filterGroup;
    const recur = (form: FormGroup | FormField): FormGroup | FormField | undefined => {
      if (form.id === id) {
        return form;
      }
      if (form.type === FormType.Group && form.children.length === 0) {
        return undefined;
      }

      if (form.type === FormType.Group) {
        for (const child of form.children) {
          const ans = recur(child);
          if (ans) {
            return ans;
          }
        }
      }
      return undefined;
    };

    const ans = recur(set);

    if (ans) {
      if (ans.type === FormType.Field) {
        if (keyType === 'columnName' && typeof value === 'string') {
          ans.columnName = value;
        } else if (keyType === 'operator' && typeof value === 'string') {
          ans.operator = value as Operator;
        } else if (keyType === 'value') {
          ans.value = value;
        }
      } else if (ans.type === FormType.Group) {
        if (keyType === 'conjunction' && (value === Conjunction.and || value === Conjunction.or)) {
          ans.conjunction = value;
        }
      }
      this.#formset.set({ filterGroup: set });
    }
  }

  public addChild(
    id: string,
    addType: FormType,
    index: number,
    obj?: Readonly<FormGroup | FormField>,
  ): void {
    const set = this.#formset.get().filterGroup;
    const recur = (form: FormGroup | FormField): void => {
      if (form.id === id && form.type === FormType.Group) {
        if (obj) {
          form.children.splice(index, 0, obj);
        } else {
          form.children.push(addType === FormType.Group ? getInitGroup() : getInitField());
        }
        return;
      }

      if (form.type === FormType.Group) {
        for (const child of form.children) {
          recur(child);
        }
      }
    };

    recur(set);
    this.#formset.set({ filterGroup: set });
  }

  public removeChild(id: string): void {
    const set = this.#formset.get().filterGroup;

    if (set.id === id) {
      this.#formset.set(structuredClone(INIT_FORMSET));
      return;
    }

    const recur = (form: FormGroup | FormField): void => {
      if (form.type === FormType.Group) {
        form.children = form.children.filter((c) => c.id !== id);
        for (const child of form.children) {
          recur(child);
        }
      }
    };
    recur(set);
    this.#formset.set({ filterGroup: set });
  }
}

export const formSets: FilterFormSet = {
  filterGroup: {
    children: [
      {
        columnName: 'tags',
        id: 'level1',
        operator: 'contains',
        type: FormType.Field,
        value: 'test',
      },
      {
        children: [
          {
            columnName: 'user',
            id: 'stringdsdff123',
            operator: 'eq',
            type: FormType.Field,
            value: 1,
          },
          {
            columnName: 'state',
            id: 'stringdsdff3',
            operator: 'notEmpty',
            type: FormType.Field,
            value: 'test',
          },
        ],
        conjunction: 'or',
        id: 'sdsdff',
        type: FormType.Group,
      },
      {
        columnName: 'name',
        id: 'stringdf123',
        operator: 'contains',
        type: FormType.Field,
        value: 'test',
      },
      {
        columnName: 'id',
        id: 'gsstringdfs123',
        operator: 'greaterEq',
        type: FormType.Field,
        value: 'test',
      },
    ],
    conjunction: Conjunction.and,
    id: 'ROOT',
    type: FormType.Group,
  },
};
