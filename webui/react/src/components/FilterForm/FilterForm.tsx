import { useObservable } from 'micro-observables';
import { DropTargetMonitor, useDrag, useDrop } from 'react-dnd';

import css from './FilterForm.module.scss';
import { FormClassStore } from './FilterFormStore';
import { FilterForm, FormField, FormGroup } from './type';

///

interface FilterFieldProps {
  field: FormField;
  conjunction: 'and' | 'or';
}

const formSets: FilterForm = {
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

const formClass = new FormClassStore(formSets);
// formClass.setValue('level1', 'TEST');
// formClass.addChild('sdf', 'field');
// formClass.addChild('gsstringdfs123', 'field');
// formClass.removeChild('gsstringdfs123');
// formClass.removeChild('sdsdff');

const ItemTypes = {
  FIELD: 'field',
  GROUP: 'group',
};
const FilterField = ({ field, conjunction }: FilterFieldProps): JSX.Element => {
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
      <button onClick={() => formClass.removeChild(field.id)}>- remove</button>
    </div>
  );
};

interface FilterGroupProps {
  group: FormGroup;
  level: number;
}

const FilterGroup = ({ group, level }: FilterGroupProps): JSX.Element => {
  const [, drag] = useDrag<FormGroup, unknown, unknown>(() => ({
    item: group,
    type: ItemTypes.GROUP,
  }));

  const [{ isOverCurrent }, drop] = useDrop<
    FormGroup | FormField,
    unknown,
    { isOverCurrent: boolean }
  >(
    () => ({
      accept: [ItemTypes.GROUP, ItemTypes.FIELD],
      collect: (monitor) => ({
        isOverCurrent: monitor.isOver({ shallow: true }),
      }),
      drop(item: FormGroup | FormField, monitor: DropTargetMonitor) {
        const isOverCurrent = monitor.isOver({ shallow: true });
        const didDrop = monitor.didDrop();
        if (isOverCurrent) {
          formClass.removeChild(item.id);
          formClass.addChild(group.id, item.type, item);
        }
        if (didDrop) {
          return;
        }
      },
    }),
    [],
  );

  const classes = [css.group];

  if (isOverCurrent) {
    classes.push(css.dragGroup);
  }

  return (
    <div className={classes.join(' ')} ref={(node) => drag(drop(node))}>
      <div>
        <div>
          {group.type} {group.id}
        </div>
        {group.children.map((child) => {
          if (child.type === 'group') {
            return <FilterGroup group={child} key={child.id} level={level + 1} />;
          } else {
            return <FilterField conjunction={group.conjunction} field={child} key={child.id} />;
          }
        })}
      </div>
      <button onClick={() => formClass.addChild(group.id, 'field')}>+ add</button>
      {level !== 1 && (
        <>
          <button onClick={() => formClass.removeChild(group.id)}>- remove</button>
        </>
      )}
    </div>
  );
};

const FilterForm = (): JSX.Element => {
  const data = useObservable(formClass.formset);

  return (
    <div>
      <FilterGroup group={data.filterSet} level={1} />
    </div>
  );
};

export default FilterForm;
