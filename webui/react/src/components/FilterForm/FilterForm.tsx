import { useObservable } from 'micro-observables';

import { FormClassStore } from './FilterFormStore';
import FilterGroup from './FilterGroup';

interface Props {
  formClassStore: FormClassStore;
}

const FilterForm = ({ formClassStore }: Props): JSX.Element => {
  const data = useObservable(formClassStore.formset);

  return (
    <div>
      <FilterGroup formClassStore={formClassStore} group={data.filterSet} level={1} />
    </div>
  );
};

export default FilterForm;
