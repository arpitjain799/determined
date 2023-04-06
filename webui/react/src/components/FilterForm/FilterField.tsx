import React from 'react';
import { useDrag } from 'react-dnd';

import css from './FilterForm.module.scss';
import { FormClassStore } from './FilterFormStore';
import { Conjunction, FormField, ItemTypes } from './type';

interface Props {
  field: FormField;
  conjunction: Conjunction;
  formClassStore: FormClassStore;
}

const FilterField = ({ field, conjunction, formClassStore }: Props): JSX.Element => {
  const [, drag] = useDrag<FormField, unknown, unknown>(() => ({
    item: field,
    type: ItemTypes.FIELD,
  }));

  return (
    <div className={css.field} ref={drag}>
      <div className={css.conjunction}>{conjunction}</div>
      <div>{field.type}</div>
      <div>{field.id}</div>
      <div>{field.columnName}</div>
      <div>{field.operator}</div>
      <div>{field.value}</div>
      <button onClick={() => formClassStore.removeChild(field.id)}>- remove</button>
    </div>
  );
};

export default FilterField;
