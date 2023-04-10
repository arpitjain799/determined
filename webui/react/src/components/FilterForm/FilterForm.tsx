import { useObservable } from 'micro-observables';

import { FilterFormStore } from './FilterFormStore';
import FilterGroup from './FilterGroup';

interface Props {
  formStore: FilterFormStore;
}

const FilterForm = ({ formStore }: Props): JSX.Element => {
  const data = useObservable(formStore.formset);
  return (
    <div>
      <FilterGroup
        conjunction={data.filterGroup.conjunction}
        formStore={formStore}
        group={data.filterGroup}
        index={0}
        level={0}
        parentId={data.filterGroup.id}
      />
      <div>{formStore.query}</div>
    </div>
  );
};

export default FilterForm;
