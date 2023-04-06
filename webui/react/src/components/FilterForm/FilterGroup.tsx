import { DropTargetMonitor, useDrag, useDrop } from 'react-dnd';

import FilterField from './FilterField';
import css from './FilterForm.module.scss';
import { FormClassStore } from './FilterFormStore';
import { FormField, FormGroup, ItemTypes } from './type';

interface Props {
  group: FormGroup;
  level: number;
  formClassStore: FormClassStore;
}

const FilterGroup = ({ group, level, formClassStore }: Props): JSX.Element => {
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
          formClassStore.removeChild(item.id);
          formClassStore.addChild(group.id, item.type, item);
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
            return (
              <FilterGroup
                formClassStore={formClassStore}
                group={child}
                key={child.id}
                level={level + 1}
              />
            );
          } else {
            return (
              <FilterField
                conjunction={group.conjunction}
                field={child}
                formClassStore={formClassStore}
                key={child.id}
              />
            );
          }
        })}
      </div>
      <button onClick={() => formClassStore.addChild(group.id, 'field')}>+ add</button>
      {level !== 1 && (
        <>
          <button onClick={() => formClassStore.removeChild(group.id)}>- remove</button>
        </>
      )}
    </div>
  );
};

export default FilterGroup;
