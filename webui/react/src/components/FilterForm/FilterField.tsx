import { DeleteOutlined, HolderOutlined } from '@ant-design/icons';
import { Select } from 'antd';
import { useDrag, useDrop } from 'react-dnd';

import Button from 'components/kit/Button';
import Input from 'components/kit/Input';
import InputNumber from 'components/kit/InputNumber';

import css from './FilterField.module.scss';
import { FilterFormStore } from './FilterFormStore';
import {
  AvaliableOperators,
  ColumnType,
  Conjunction,
  FormField,
  FormGroup,
  FormType,
  Operator,
} from './type';

interface Props {
  index: number; // start from 0
  field: FormField;
  parentId: string;
  conjunction: Conjunction;
  formStore: FilterFormStore;
  level: number; // start from 0
}

const FilterField = ({
  field,
  conjunction,
  formStore,
  index,
  parentId,
  level,
}: Props): JSX.Element => {
  const [, drag, preview] = useDrag<{ form: FormField; index: number }, unknown, unknown>(() => ({
    item: { form: field, index },
    type: FormType.Field,
  }));

  const [{ isOverCurrent, canDrop }, drop] = useDrop<
    { form: FormGroup | FormField; index: number },
    unknown,
    { isOverCurrent: boolean; canDrop: boolean }
  >({
    accept: [FormType.Group, FormType.Field],
    canDrop(item, monitor) {
      const isOverCurrent = monitor.isOver({ shallow: true });
      if (isOverCurrent) {
        if (item.form.type === FormType.Group) {
          return (
            // cant dnd with deeper than 2 level group
            level < 2 &&
            // cant dnd if sum of source children of group type (0 if none, 1 if children exist)
            // and target item's level is over 3 for field
            (item.form.children.filter((c) => c.type === FormType.Group).length === 0 ? 0 : 1) +
              level <
              3
          );
        }
        return true;
      }
      return false;
    },
    collect: (monitor) => ({
      canDrop: monitor.canDrop(),
      isOverCurrent: monitor.isOver({ shallow: true }),
    }),
    hover(item) {
      const dragIndex = item.index;
      const hoverIndex = index;
      if (dragIndex !== hoverIndex && isOverCurrent && canDrop) {
        formStore.removeChild(item.form.id);
        formStore.addChild(parentId, item.form.type, hoverIndex, item.form);
        item.index = hoverIndex;
      }
    },
  });

  return (
    <div className={css.base} ref={(node) => drop(node)}>
      <>
        {index === 0 && <div>if</div>}
        {index === 1 && (
          <Select
            value={conjunction}
            onChange={(value: string) => {
              formStore.setFieldValue(parentId, 'conjunction', value);
            }}>
            {Object.values(Conjunction).map((c) => (
              <Select.Option key={c} value={c}>
                {c}
              </Select.Option>
            ))}
          </Select>
        )}
        {index > 1 && <div className={css.conjunction}>{conjunction}</div>}
      </>
      <div className={css.fieldCard} ref={preview}>
        <Select
          value={field.columnName}
          onChange={(value: string) => {
            const prevType = ColumnType[field.columnName];
            formStore.setFieldValue(field.id, 'columnName', value);
            if (prevType !== ColumnType[field.columnName]) {
              const defaultOperator = AvaliableOperators[ColumnType[field.columnName]][0];
              formStore.setFieldValue(field.id, 'operator', defaultOperator);
              formStore.setFieldValue(field.id, 'value', null);
            }
          }}>
          {Object.keys(ColumnType).map((col) => (
            <Select.Option key={col} value={col}>
              {col}
            </Select.Option>
          ))}
        </Select>
        <Select
          style={{ width: '100%' }}
          value={field.operator}
          onChange={(value: Operator) => {
            formStore.setFieldValue(field.id, 'operator', value);
          }}>
          {AvaliableOperators[ColumnType[field.columnName]].map((op) => (
            <Select.Option key={op} value={op}>
              {op}
            </Select.Option>
          ))}
        </Select>
        <>
          {ColumnType[field.columnName] === 'string' && (
            <Input
              disabled={field.operator === Operator.isEmpty || field.operator === Operator.notEmpty}
              value={
                field.operator === Operator.isEmpty || field.operator === Operator.notEmpty
                  ? undefined
                  : field.value?.toString()
              }
              onChange={(e) => formStore.setFieldValue(field.id, 'value', e.target.value)}
            />
          )}
          {ColumnType[field.columnName] === 'number' && (
            <InputNumber
              value={field.value != null ? Number(field.value) : undefined}
              onChange={(val) =>
                formStore.setFieldValue(field.id, 'value', val?.toString() ?? null)
              }
            />
          )}
        </>
        <Button
          icon={<DeleteOutlined />}
          type="text"
          onClick={() => formStore.removeChild(field.id)}
        />
        <Button type="text">
          <div ref={drag}>
            <HolderOutlined />
          </div>
        </Button>
      </div>
    </div>
  );
};

export default FilterField;
