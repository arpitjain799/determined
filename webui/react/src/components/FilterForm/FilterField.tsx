import { DeleteOutlined, HolderOutlined } from '@ant-design/icons';
import { Select } from 'antd';
import { useDrag } from 'react-dnd';

import Button from 'components/kit/Button';
import Input from 'components/kit/Input';
import InputNumber from 'components/kit/InputNumber';

import css from './FilterField.module.scss';
import { FormClassStore } from './FilterFormStore';
import { Conjunction, FormField, ItemTypes, Operator, OperatorMap } from './type';

interface Props {
  index: number; // start from 0
  field: FormField;
  conjunction: Conjunction;
  formClassStore: FormClassStore;
}

const FilterField = ({ field, conjunction, formClassStore, index }: Props): JSX.Element => {
  const [, drag, preview] = useDrag<FormField, unknown, unknown>(() => ({
    item: field,
    type: ItemTypes.FIELD,
  }));

  return (
    <div className={css.base}>
      {index !== 0 ? <div>{conjunction}</div> : <div>where</div>}
      <div className={css.fieldCard} ref={preview}>
        <div>{field.columnName}</div>
        <Select
          style={{ width: '150px' }}
          value={field.operator}
          onChange={(value: Operator) => {
            formClassStore.setFieldValue(field.id, 'operator', value);
          }}>
          {Object.entries(OperatorMap).map((op) => (
            <Select.Option key={op[0]} value={op[0]}>
              {op[1]}
            </Select.Option>
          ))}
        </Select>
        {['string'].includes(field.columnName) ? (
          <Input size="small" value={field.value?.toString()} />
        ) : (
          <InputNumber value={field.value as number} />
        )}
        <Button
          icon={<DeleteOutlined />}
          type="text"
          onClick={() => formClassStore.removeChild(field.id)}
        />
        <div className={css.draggableHandle} ref={drag}>
          <Button type="text">
            <HolderOutlined />
          </Button>
        </div>
      </div>
    </div>
  );
};

export default FilterField;
