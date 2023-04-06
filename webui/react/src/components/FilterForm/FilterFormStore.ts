import { Observable, observable } from 'micro-observables';

import { FilterFormSet, FormField, FormGroup } from './type';

export class FormClassStore {
  #formset = observable<FilterFormSet>({
    filterSet: { children: [], conjunction: 'and', id: 'ROOT', type: 'group' }, // default
  });

  constructor(data?: FilterFormSet) {
    if (data) {
      this.#formset = observable<FilterFormSet>(data);
    }
  }

  public get formset(): Observable<Readonly<FilterFormSet>> {
    return this.#formset.readOnly();
  }

  public setValue(id: string, value: string | string[] | number | number[]): void {
    const set = this.#formset.get().filterSet;
    const recur = (form: FormGroup | FormField): FormGroup | FormField | undefined => {
      if (form.id === id) {
        return form;
      }
      if (form.type === 'group' && form.children.length === 0) {
        return undefined;
      }

      if (form.type === 'group') {
        for (const child of form.children) {
          const ans = recur(child);
          if (ans) {
            return ans;
          }
        }
      }
    };
    const ans = recur(set);

    if (ans) {
      if (ans.type === 'field') {
        ans.value = value;
        this.#formset.set({ filterSet: set });
      }
    }
  }

  public addChild(id: string, addType: 'group' | 'field', obj?: FormGroup | FormField): void {
    const set = this.#formset.get().filterSet;
    const recur = (form: FormGroup | FormField): void => {
      if (form.id === id && form.type === 'group') {
        if (obj) {
          form.children.push({ ...obj });
          return;
        }
        if (addType === 'group') {
          form.children.push({
            children: [],
            conjunction: 'and',
            id: 'sdf',
            type: 'group',
          });
        } else {
          form.children.push({
            columnName: 'radfdsf',
            id: crypto.randomUUID(),
            operator: 'contains',
            type: 'field',
            value: 'new',
          });
        }
        return;
      }
      if (form.type === 'group' && form.children.length === 0) {
        return;
      }

      if (form.type === 'group') {
        for (const child of form.children) {
          recur(child);
        }
      }
    };
    recur(set);
    this.#formset.set({ filterSet: set });
  }

  public removeChild(id: string): void {
    const set = this.#formset.get().filterSet;
    const recur = (form: FormGroup | FormField): void => {
      if (form.type === 'group') {
        form.children = form.children.filter((c) => c.id !== id);
        for (const child of form.children) {
          recur(child);
        }
      }
    };
    recur(set);
    this.#formset.set({ filterSet: set });
  }
}

export const formSets: FilterFormSet = {
  filterSet: {
    children: [
      {
        columnName: 'level1',
        id: 'level1',
        operator: 'contains',
        type: 'field',
        value: 'test',
      },
      {
        children: [
          {
            columnName: 'string',
            id: 'stringdsdff123',
            operator: 'contains',
            type: 'field',
            value: 'test',
          },
          {
            columnName: 'string',
            id: 'stringdsdff3',
            operator: 'contains',
            type: 'field',
            value: 'test',
          },
        ],
        conjunction: 'and',
        id: 'sdsdff',
        type: 'group',
      },
      {
        columnName: 'string',
        id: 'stringdf123',
        operator: 'contains',
        type: 'field',
        value: 'test',
      },
      {
        columnName: 'string',
        id: 'gsstringdfs123',
        operator: 'contains',
        type: 'field',
        value: 'test',
      },
    ],
    conjunction: 'and',
    id: 'ROOT',
    type: 'group',
  },
};
