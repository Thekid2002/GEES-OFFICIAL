import type { AttributeType } from '@/models/attribute-type.ts'

export class AttributeDisplayType {
  displayName: string;
  type: AttributeType;
  attributeName: string;
  icon: string | null;
  editable: boolean;
  description: string | null;

  constructor(
    displayName: string,
    type: AttributeType,
    attributeName: string,
    icon: string | null = null,
    description: string | null = null,
    editable: boolean = false
  ) {
    this.displayName = displayName;
    this.type = type;
    this.attributeName = attributeName;
    this.icon = icon;
    this.editable = editable;
    this.description = description;
  }
}
