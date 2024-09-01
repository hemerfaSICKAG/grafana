import { RegistryItem } from '@grafana/data';
import { SceneObject, VizPanel } from '@grafana/scenes';
import { OptionsPaneItemDescriptor } from 'app/features/dashboard/components/PanelEditor/OptionsPaneItemDescriptor';

export interface DashboardLayoutManager extends SceneObject {
  /**
   * Notify the layout manager that the edit mode has changed
   * @param isEditing
   */
  editModeChanged(isEditing: boolean): void;
  /**
   * We should be able to figure out how to add the explore panel in a way that leaves the
   * initialSaveModel clean from it so we can leverage the default discard changes logic.
   * Then we can get rid of this.
   */
  cleanUpStateFromExplore?(): void;
  /**
   * Not sure we will need this in the long run, we should be able to handle this inside internally
   */
  getNextPanelId(): number;
  /**
   * Used for transferring state between layouts. Not sure what the return type should be here.
   * Right now we just check for VizPanels
   */
  getElements(): LayoutElementInfo[];
  /**
   * Renders options and layout actions
   */
  renderEditor?(): React.ReactNode;
  /**
   * Get's the layout descriptor (which has the name and id)
   */
  getDescriptor(): LayoutRegistryItem;
  /**
   * When switching between layouts
   * @param currentLayout
   */
  initFromLayout(currentLayout: DashboardLayoutManager): DashboardLayoutManager;
  /**
   * Create from persisted state
   * @param saveModel
   */
  initFromSaveModel?(saveModel: any): void;
  /**
   * Turn into a save model
   * @param saveModel
   */
  toSaveModel?(): any;
}

/**
 * The layout descriptor used when selecting / switching layouts
 */
export interface LayoutRegistryItem extends RegistryItem {
  /**
   * This is for creating a new layout from the elements of another layout
   * @param elements
   * @returns
   */
  create: () => DashboardLayoutManager;
}

export interface LayoutEditorProps<T> {
  layoutManager: T;
}

export interface LayoutElementInfo {
  body: SceneObject;
  width?: number;
  height?: number;
}

export interface LayoutParent extends SceneObject {
  switchLayout(newLayout: DashboardLayoutManager): void;
}

export function isLayoutParent(obj: SceneObject): obj is LayoutParent {
  return 'switchLayout' in obj;
}

export interface PanelEditLayoutElement {
  getOptions(): OptionsPaneItemDescriptor[];
}
