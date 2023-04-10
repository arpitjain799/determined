import { FilterOutlined } from '@ant-design/icons';
import { Button, Popconfirm } from 'antd';

import FilterForm from 'components/FilterForm/FilterForm';
import { FilterFormStore, formSets } from 'components/FilterForm/FilterFormStore';

const formStore = new FilterFormStore(formSets);

const TEST = (): JSX.Element => {
  return (
    <>
      <FilterForm formStore={formStore} />
      <div style={{ display: 'flex', justifyContent: 'center', marginTop: '40px' }}>
        <Popconfirm
          description={<FilterForm formStore={formStore} />}
          icon={<FilterOutlined />}
          title={'Table Filter'}>
          <Button>Click ME</Button>
        </Popconfirm>
      </div>
    </>
  );
};

export default TEST;
