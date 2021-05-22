/*
 * lakeFS API
 * lakeFS HTTP API
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package io.lakefs.clients.api.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * UnderlyingObjectProperties
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class UnderlyingObjectProperties {
  public static final String SERIALIZED_NAME_STORAGE_CLASS = "storage_class";
  @SerializedName(SERIALIZED_NAME_STORAGE_CLASS)
  private String storageClass;


  public UnderlyingObjectProperties storageClass(String storageClass) {
    
    this.storageClass = storageClass;
    return this;
  }

   /**
   * Get storageClass
   * @return storageClass
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStorageClass() {
    return storageClass;
  }


  public void setStorageClass(String storageClass) {
    this.storageClass = storageClass;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UnderlyingObjectProperties underlyingObjectProperties = (UnderlyingObjectProperties) o;
    return Objects.equals(this.storageClass, underlyingObjectProperties.storageClass);
  }

  @Override
  public int hashCode() {
    return Objects.hash(storageClass);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UnderlyingObjectProperties {\n");
    sb.append("    storageClass: ").append(toIndentedString(storageClass)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

